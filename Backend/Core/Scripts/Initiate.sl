set Frontend := import("Backend/Core/Scripts/Frontend.sl");

set Interface := "";
set ARPState := "";
set SSDPState := "";
set TimerArp := 5; // 5 seconds is standard | converted to int on casters backend in the plugin

define init() {
    ENGINE("Backend/Conf/EngineFiles/Engine.slmod")
    set Arguments := args();
    io.clear()
    if (isroot() == "1000") {
        Frontend::Banner()
        set expectedmatch := int(2);
        if ( Arguments.Length() > int(1)) {
            Arguments.PopL() // Lets get that file name out of the argument list
            foreach arg in Arguments {
                // split each argument, lets go ahead and get its name and the setting
                set splitter := arg.Split("=");
                if (splitter.Length() >= expectedmatch) {
                    switch (splitter[0].Lower()) {
                        case "interface" {
                            Frontend::Logger("Interface Being Used -> [{}]".Format(splitter[1]), "debug")
                            set net := network.NetByName(splitter[1]);
                            if (net == "Unknown") {
                                Frontend::Logger("The interface provided [{}{}{}] did not have any plausible addresses".Format(
                                    console.HtmlToAnsi("56fc03", ""),
                                    splitter[0], 
                                    console.HtmlToAnsi("0ee6c2", "")
                                ), "debug")
                                Frontend::Logger("This means that Caster's module will attempt to find a interface that is stable", "debug")
                            } else {
                                Assign("Interface", splitter[1])
                            }
                        }
                        case "delim" {
                            Assign("TimerArp", splitter[1].ToInt())
                        }
                        case "arp" {
                            //Frontend::Logger("ARP Module is -> {}".Format(splitter[1]), "debug")
                            Assign("ARPState", splitter[1])
                        }
                        case "ssdp" {
                            //Frontend::Logger("SSDP Module is -> {}".Format(splitter[1]), "debug")
                            Assign("SSDPState", splitter[1])
                        }
                        default {
                            //Frontend::Logger("BAD MODULE NAME [{}] WITH VALUE [{}]".Format(splitter[0], splitter[1]))
                        }
                    }
                }
                if (arg == "help") {
                    // Display help
                    Frontend::Help()
                    quit(int(0))
                }
            } 
            // Now we can go ahead and import the plugin and pass the data that Caster's backend module needs. The plugin will need some work but its not that bad
            InitiatePlugin("Backend/Core/BinPlugins/Caster");
            if (Interface == "") {
                Frontend::Logger("Plugin is loading interfaces...", "debug");
                ExecuteFF("Framework", "StartArp", false, "");
            } else {
                Frontend::Logger("Plugin is using [{}] for Casters ARP module...".Format(Interface));
                ExecuteFF("Framework", "StartArp", true, Interface);
            }
            if (PluginState()) {
                Frontend::Logger("Plugin state is [{}Active{}]!".Format(console.HtmlToAnsi("56fc03", ""), console.HtmlToAnsi("0ee6c2", "")), "debug")
            } else {
                Frontend::Logger("Plugin state is invalid...this is a bad thing- horrid thing actually...", "error")
                quit(int(0))
            }
            ExecuteFF("Framework", "WriteRoutine", "Backend/UnitDatabase/Capture.json", TimerArp)
            Frontend::RestoreStd() 
            Frontend::NewLn()
        } else {
            Frontend::Logger("At least specify an interface or type 'help' for more", "error")
            quit(int(0));
        }
    } else {
        println("[Caster] Error: You need to be root to work this framework")
        quit(int(0))
    }
}


// For some reason, removing ':=' from this will cause the error `io.input identifier not found` which is fucking weird
set ExportableSettings := {
    "ssdp": SSDPState,
    "arp": ARPState,
    "interface": Interface
}; 