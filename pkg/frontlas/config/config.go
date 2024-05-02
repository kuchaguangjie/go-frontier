package config

import (
	"flag"
	"fmt"
	"io"
	"os"

	armio "github.com/jumboframes/armorigo/io"
	"github.com/jumboframes/armorigo/log"
	"github.com/singchia/frontier/pkg/config"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"
)

// daemon related
type RLimit struct {
	Enable  bool `yaml:"enable"`
	NumFile int  `yaml:"nofile"`
}

type PProf struct {
	Enable         bool   `yaml:"enable"`
	Addr           string `yaml:"addr"`
	CPUProfileRate int    `yaml:"cpu_profile_rate"`
}

type Daemon struct {
	RLimit RLimit `yaml:"rlimit"`
	PProf  PProf  `yaml:"pprof"`
}

// for rest and grpc
type ControlPlane struct {
	Listen config.Listen `yaml:"listen"`
}

// TODO tls support
type Redis struct {
	Mode       string `yaml:"mode"` // standalone, sentinel or cluster
	Standalone struct {
		// The network type, either tcp or unix.
		// Default is tcp.
		Network string
		// host:port address.
		Addr string
		// Protocol 2 or 3. Use the version to negotiate RESP version with redis-server.
		// Default is 3.
		Protocol int
		// Use the specified Username to authenticate the current connection
		// with one of the connections defined in the ACL list when connecting
		// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
		Username string
		// Optional password. Must match the password specified in the
		// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
		// or the User Password when connecting to a Redis 6.0 instance, or greater,
		// that is using the Redis ACL system.
		Password string
		// CredentialsProvider allows the username and password to be updated
		// before reconnecting. It should return the current username and password.
		DB int
		// ClientName will execute the `CLIENT SETNAME ClientName` command for each conn.
		ClientName string

		// connection retry settings
		MaxRetries      int
		MinRetryBackoff int
		MaxRetryBackoff int

		// connection r/w settings
		DialTimeout  int
		ReadTimeout  int
		WriteTimeout int

		// connection pool settings
		PoolFIFO        bool
		PoolSize        int // applies per cluster node and not for the whole cluster
		PoolTimeout     int
		MinIdleConns    int
		MaxIdleConns    int
		MaxActiveConns  int // applies per cluster node and not for the whole cluster
		ConnMaxIdleTime int
		ConnMaxLifetime int

		DisableIndentity bool   // Disable set-lib on connect. Default is false.
		IdentitySuffix   string // Add suffix to client name. Default is empty.
	}
	Sentinel struct {
		Addrs      []string `yaml:"addrs"`
		MasterName string   `yaml:"master_name"`
		Protocol   int
		Username   string
		Password   string
		DB         int
		// ClientName will execute the `CLIENT SETNAME ClientName` command for each conn.
		ClientName string

		// route settings
		// Allows routing read-only commands to the closest master or replica node.
		// This option only works with NewFailoverClusterClient.
		RouteByLatency bool
		// Allows routing read-only commands to the random master or replica node.
		// This option only works with NewFailoverClusterClient.
		RouteRandomly bool
		// Route all commands to replica read-only nodes.
		ReplicaOnly bool

		// Use replicas disconnected with master when cannot get connected replicas
		// Now, this option only works in RandomReplicaAddr function.
		UseDisconnectedReplicas bool

		// connection retry settings
		MaxRetries      int
		MinRetryBackoff int
		MaxRetryBackoff int

		// connection r/w settings
		DialTimeout  int
		ReadTimeout  int
		WriteTimeout int

		// connection pool settings
		PoolFIFO        bool
		PoolSize        int // applies per cluster node and not for the whole cluster
		PoolTimeout     int
		MinIdleConns    int
		MaxIdleConns    int
		MaxActiveConns  int // applies per cluster node and not for the whole cluster
		ConnMaxIdleTime int
		ConnMaxLifetime int

		DisableIndentity bool   // Disable set-lib on connect. Default is false.
		IdentitySuffix   string // Add suffix to client name. Default is empty.
	}
	Cluster struct {
		Addrs    []string `yaml:"addrs"`
		Protocol int
		Username string
		Password string

		// ClientName will execute the `CLIENT SETNAME ClientName` command for each conn.
		ClientName string

		// The maximum number of retries before giving up. Command is retried
		// on network errors and MOVED/ASK redirects.
		// Default is 3 retries.
		MaxRedirects int

		// Allows routing read-only commands to the closest master or slave node.
		// It automatically enables ReadOnly.
		RouteByLatency bool
		// Allows routing read-only commands to the random master or slave node.
		// It automatically enables ReadOnly.
		RouteRandomly bool

		// connection retry settings
		MaxRetries      int
		MinRetryBackoff int
		MaxRetryBackoff int

		// connection r/w settings
		DialTimeout  int
		ReadTimeout  int
		WriteTimeout int

		// connection pool settings
		PoolFIFO        bool
		PoolSize        int // applies per cluster node and not for the whole cluster
		PoolTimeout     int
		MinIdleConns    int
		MaxIdleConns    int
		MaxActiveConns  int // applies per cluster node and not for the whole cluster
		ConnMaxIdleTime int
		ConnMaxLifetime int

		DisableIndentity bool   // Disable set-lib on connect. Default is false.
		IdentitySuffix   string // Add suffix to client name. Default is empty.
	}
}

type FrontierManager struct {
	Listen config.Listen `yaml:"listen"`
}

type Configuration struct {
	Daemon Daemon `yaml:"daemon"`

	ControlPlane ControlPlane `yaml:"control_plane"`

	FrontierManager FrontierManager `yaml:"frontier_plane"`

	Redis Redis `yaml:"redis"`
}

func Parse() (*Configuration, error) {
	var (
		argConfigFile         = pflag.String("config", "", "config file, default not configured")
		argArmorigoLogLevel   = pflag.String("loglevel", "info", "log level for armorigo log")
		argDaemonRLimitNofile = pflag.Int("daemon-rlimit-nofile", -1, "SetRLimit for number of file of this daemon, default: -1 means ignore")
		// TODO more command-line args

		config *Configuration
	)
	pflag.Lookup("daemon-rlimit-nofile").NoOptDefVal = "1048576"

	// set klog
	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)

	// sync the glog and klog flags.
	pflag.CommandLine.VisitAll(func(f1 *pflag.Flag) {
		f2 := klogFlags.Lookup(f1.Name)
		if f2 != nil {
			value := f1.Value.String()
			if err := f2.Value.Set(value); err != nil {
				klog.Fatal(err, "failed to set flag")
				return
			}
		}
	})

	pflag.CommandLine.AddGoFlagSet(klogFlags)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	// armorigo log
	level, err := log.ParseLevel(*argArmorigoLogLevel)
	if err != nil {
		fmt.Println("parse log level err:", err)
		return nil, err
	}
	log.SetLevel(level)
	log.SetOutput(os.Stdout)

	// config file
	if *argConfigFile != "" {
		// TODO the command-line is prior to config file
		data, err := os.ReadFile(*argConfigFile)
		if err != nil {
			return nil, err
		}
		config = &Configuration{}
		if err = yaml.Unmarshal(data, config); err != nil {
			return nil, err
		}
	}

	if config == nil {
		config = &Configuration{}
	}
	// daemon
	config.Daemon.RLimit.NumFile = *argDaemonRLimitNofile
	if config.Daemon.PProf.CPUProfileRate == 0 {
		config.Daemon.PProf.CPUProfileRate = 10000
	}
	return config, nil
}

func genDefaultConfig(writer io.Writer) error {
	conf := &Configuration{
		Daemon: Daemon{
			RLimit: RLimit{
				NumFile: 1024,
			},
			PProf: PProf{
				Enable: true,
				Addr:   "0.0.0.0:6061",
			},
		},
		ControlPlane: ControlPlane{
			Listen: config.Listen{
				Network: "tcp",
				Addr:    "0.0.0.0:30020",
			},
		},
		FrontierManager: FrontierManager{
			Listen: config.Listen{
				Network: "tcp",
				Addr:    "0.0.0.0:30021",
			},
		},
		Redis: Redis{},
	}
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	_, err = armio.WriteAll(data, writer)
	if err != nil {
		return err
	}
	return nil
}
