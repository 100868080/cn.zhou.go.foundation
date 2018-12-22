if err := WaitForServer(url); err != nil {
	fmt.Fprintf(os.Stderr, "Site is down:%v\n", err)
	os.Exit(1)
}
if err := WaitForServer(url); err != nil {
	log.Fatal("Site is down:%v\n", err)
}

log.SetPrefix("wait:")
log.SetFlags(0)

if err := Ping(); err != nil {
	log.Printf("ping failed:%v;networking disabled", err)
}

if err := Ping(); err != nil {
	fmt.Fprintf(os.Stderr, "ping failed:%v;networking disabled\n", err)
}

dir, err := ioutil.TempDir("", "scratch")
if err != nil {
	return fmt.Errorf("failed to create temp dir:%v", err)
}

os.RemoveAl(dir)