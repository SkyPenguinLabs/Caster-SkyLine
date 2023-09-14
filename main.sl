set CasterCore := import("Backend/Core/Scripts/Initiate.sl");
set CasterHandler := import("Backend/Core/Scripts/DataLoader.sl");
set CasterPluginFuncs := import("Backend/Core/Scripts/PluginBuffers.sl");
set Verification := import("Backend/Core/Scripts/Verification.sl");
set Requesting := import("Backend/Core/Scripts/Requests.sl");
set CurrentLocation := "local";


define Main() {
    // main entry point
    for(true) {
        set CasterIn := io.input("Caster@{}> ".Format(CurrentLocation), "n");
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
                    default {
                        println("Tabulating default")
                        CasterHandler::TabulateHosts("all")
                    }
                }
            } else {
                println("What are you trying to view?")
            }

        } else if (CasterIn.Index(int(4)) == "help") {
            println("helo!")
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
                        if (EnumerateCommand.Contains("apple")) {
                            if (DeviceDB == "apple") {
                                Requesting::EnumerateDevice(EnumerateCommand, Target)
                            } else {
                                println("Error: You can not target an apple device on this address. The device must be an Apple device")
                            }
                        } else if (EnumerateCommand.Contains("roku")) {
                            if (DeviceDB == "roku") {

                            } else {

                            }
                        } else if (EnumerateCommand.Contains("google")) {
                            if (DeviceDB == "google") {

                            } else {

                            }
                        } else if (EnumerateCommand.Contains("amazon")) {
                            if (DeviceDB == "amazon") {
                                 
                            } else {

                            }
                        }
                        println(newarr[1].Pop(" ").Lower())
                        Requesting::EnumerateDevice(Res[0]["IP"], newarr[1].Pop(" ").Lower())
                    }
                } else {
                    println("What are you trying to enumerate?")
                }
            }
        }
    }
}; 

Main()
