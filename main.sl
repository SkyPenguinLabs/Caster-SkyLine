set CasterCore := import("Backend/Core/Scripts/Initiate.sky");
set CasterHandler := import("Backend/Core/Scripts/DataLoader.sky");
set CasterPluginFuncs := import("Backend/Core/Scripts/PluginBuffers.sky");
set Verification := import("Backend/Core/Scripts/Verification.sky");
set Requesting := import("Backend/Core/Scripts/Requests.sky");
set CurrentLocation := "local";


define Main() {
    // main entry point
    if (CasterCore::CLI_Mode == true) {
        if (CasterCore::CLI_Session_String_Binary_File == true) {
            if (CasterCore::CLI_Session_File_For_File_Operations != "") {
                mode("pwn");
                set res := Strings(CasterCore::CLI_Session_File_For_File_Operations);
                set data := "";

                foreach char in res {
                    data += char + " \n ";
                }

                println(io.FormatBox(
                    Unique.HighlightPattern(
                        data,
                        [
                            console.HtmlToAnsi("0ee6c2", ""),
                            console.HtmlToAnsi("fa5b05", "")
        
                        ],
                        [
                            `([0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12})`,
                            `([0-9a-fA-F]{2}[:-]){5}([0-9a-fA-F]{2})`
                        ]
                    )
                ))
            }
        } 
        if (CasterCore::RokuECPMode == true) {
            set Res := json.Parse("Backend/UnitDatabase/TargetInfo.json")["Device"];
            if (Res[0]["IP"] == "") {
                println("Please for the love of fucking jesus himself- set a target")
                quit(int(0))
            } else {
                set Target := Res[0]["IP"];
                ExecuteFF("Framework", "EnterECP", Target)
                quit(int(0))
            }
        }
        quit(int(0))

    } else {
        for(true) {
            set CasterIn := Requesting::CasterInputStream::CasterInput("{}Command{}".Format(
                console.HtmlToAnsi("5d42f5", ""),
                console.HtmlToAnsi("42f5bf", "")
            ));
            if (CasterIn.Index(int(3)) == "set") {    
                CasterIn = CasterIn.Pop("set") // popping set from the input
                
                if (CasterIn.Contains("=")) {
                    
                    set newarr := CasterIn.Split("=");
                    set Expected := int(2);
                    
                    if (newarr.Length() >= Expected) {
                        switch (newarr[0].Pop(" ").Lower()) {
                            case "device" {
                                Verification::VerifyHostOnNetwork(newarr[1].Pop(" ").Lower())
                            }
                        }
                    } else {
                        println("Error: Caster needs you to specify a variable and a value -> `set x=y`")
                    }

                }
            } else if (CasterIn.Index(int(4)) == "view") {
                set newarr := CasterIn.Split(" ");
                set expect := int(1);    
                if (newarr.Length() >= expect) {
                    switch (newarr[1].Pop(" ").Lower()) {
                        case "google" {
                            CasterHandler::TabulateHosts("Google")
                        }
                        case "amazon" {
                            CasterHandler::TabulateHosts("Amazon")
                        }
                        case "apple" {
                            CasterHandler::TabulateHosts("Apple")
                        }
                        case "roku" {
                            CasterHandler::TabulateHosts("Roku")
                        }
                        case "hosts" {
                            CasterHandler::TabulateHosts("all")
                        }
                        case "clear" {
                            io.clear()
                        }
                        case "clearout" {
                            io.clear()
                            CasterCore::Frontend::Banner()
                        }
                        case "settings" {
                            set Res := json.Parse("Backend/UnitDatabase/TargetInfo.json")["Device"];
                            console.TableNew(
                                console.HtmlToAnsi(
                                    "03fcb6", "", false
                                ), 
                                console.HtmlToAnsi(
                                    "ffffff", "", false
                                ),  
                                console.HtmlToAnsi(
                                    "ffffff", "", false
                                ),  
                                console.HtmlToAnsi(
                                    "ffffff", "", false
                                ), 
                                console.HtmlToAnsi(
                                    "ffffff", "", false
                                ), 
                                `━`, 
                                `┃`, 
                                `┫`, 
                                `┣`
                            );
                            set Cols := ["Device Name", "Device Database", "Device IP"];
                            set Rows := [];
                            foreach node in Res {
                                Rows.Append(
                                    [node["Brand"], node["Database"], node["IP"]]
                                )
                            };
                            println(console.Table(Rows, Cols))
                        }
                        default {
                            println("Tabulating default")
                            CasterHandler::TabulateHosts("all")
                        }
                    }
                } else {
                    println("What are you trying to view?")
                }

            } else if (CasterIn.Index(int(4)) == "help") {
                CasterCore::Frontend::ConsoleHelp()
            } else if (CasterIn.Index(int(6)) == "import") {
                CasterPluginFuncs::PluginUpdateFile();
                println("[+] Data has been updated")
            } else if (CasterIn.Index(int(9)) == "enumerate") {
                set newarr := CasterIn.Split(" ");
                set expect := int(1);
                if (newarr.Length() >= expect) {
                    // I should most likely check if a target address is set
                    if (newarr[1].Pop(" ").Lower() != "") {
                        set Res := json.Parse("Backend/UnitDatabase/TargetInfo.json")["Device"];
                        if (Res[0]["IP"] == "") {
                            println("Please for the love of fucking jesus himself- set a target")
                        } else {
                            /*
                                We now need to set some form of enumeration 
                                for the command map. Per device, so, we need to 
                                also say what we want to enumerate directly. Commands
                                should be targeting
                            
                            */
                            set Target := Res[0]["IP"];
                            set DeviceDB := Res[0]["Database"];
                            set EnumerateCommand := newarr[1].Pop(" ").Lower();
                            switch (CasterCore::Strict) {
                                case "off" {
                                    if (EnumerateCommand.Contains("apple") == true) {
                                        Requesting::EnumerateDevice(EnumerateCommand, Target)
                                    } else if (EnumerateCommand.Contains("roku") == true) {
                                        Requesting::EnumerateDevice(EnumerateCommand, Target)
                                    } else if (EnumerateCommand.Contains("amazon") == true) {
                                        Requesting::EnumerateDevice(EnumerateCommand, Target)
                                    } else if (EnumerateCommand.Contains("cast") == true) {
                                        Requesting::EnumerateDevice(EnumerateCommand, Target)
                                    } else {
                                        println("Unsupported brand")
                                    }
                                }
                                case "on" {
                                    if (EnumerateCommand.Contains("apple") == true) {
                                        if (DeviceDB == "apple") {
                                            Requesting::EnumerateDevice(EnumerateCommand, Target)
                                        } else {
                                            println("Error: You can not target an apple device on this address. The device must be an Apple device")
                                        }
                                    } else if (EnumerateCommand.Contains("roku") == true) {
                                        if (DeviceDB == "roku") {
                                            Requesting::EnumerateDevice(EnumerateCommand, Target)
                                        } else {
                                            println("Error: You can not target a Roku device on this address. The device must be an Roku device")
                                        }
                                    } else if (EnumerateCommand.Contains("google") == true) {
                                        if (DeviceDB == "google") {
                                            Requesting::EnumerateDevice(EnumerateCommand, Target)
                                        } else {

                                        }
                                    } else if (EnumerateCommand.Contains("amazon") == true) {
                                        if (DeviceDB == "amazon") {
                                            Requesting::EnumerateDevice(EnumerateCommand, Target)
                                        } else {
                                            println("Error: You can not target a Amazon device on this address. The device must be an Amazon device")
                                        }
                                    } else {
                                        println("Unsupported brand")
                                    }
                                }
                                default {
                                    println("Woah- dev error here| Unknown state has been set")
                                    quit(
                                        int(
                                            0
                                        )
                                    )
                                }
                            }
                        }
                    } else {
                        println("What are you trying to enumerate?")
                    }
                }
            }
        }
    }
}; 

Main()
