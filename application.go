func Completion() *cobra.Command {
	// completionCmd outputs the completion script
	c := &cobra.Command{
		Use:   "completions",
		Short: "Generate tab completion scripts",

    }}
type initArgs struct {
	algo              string
	chainID           string
	keyringBackend    string
	minGasPrices      string
	nodeDaemonHome    string
	nodeDirPrefix     string
	numValidators     int
	outputDir         string
	startingIPAddress string
}

func main() {
	log.SetPrefix("mkzip: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 || !strings.HasSuffix(args[0], ".zip") {
		usage()
	} 

  
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				err := cmd.Root().GenBashCompletion(os.Stdout)
				if err != nil {
					panic(err)
				}
