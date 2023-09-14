
define EnumerateDevice(command, target) {
    switch(command) {
        case "apple-serverinfo" {
            set Body := http.Get(
                sprintf(
                    Apple.AirPlayServerInfo,
                    target
                )
            );
            set XmlResponse := Body["ResponseBody"];
            println(XmlResponse)
        }
        default {
            println("Error: Command does not exist")
        }
    }
} 