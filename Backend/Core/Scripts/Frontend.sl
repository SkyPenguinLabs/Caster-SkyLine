set VarsNonMod := import("Backend/Core/Scripts/VarsLocalNonModifiable.sl");

define RestoreStd(color="56fc03") {
    println(console.HtmlToAnsi(color, ""))
}

define Banner() {
    set contents := File.Open("Backend/Core/Data/Banner.txt");
    println(
        contents.Format(
            console.HtmlToAnsi("4c02a8", ""),
            console.HtmlToAnsi("FF5F00", ""),
            console.HtmlToAnsi("FF5F00", ""),
            console.HtmlToAnsi("FF5F00", ""),
            console.HtmlToAnsi("4c02a8", "")
        )
    )
    RestoreStd()
}

define NewLn() {
    println("")
}

define Logger(message, type="") {
    switch (type) {
        case "debug" {
            print("{}[Caster] {}DEBUG: {}".Format(
                console.HtmlToAnsi("fa5b05", ""),
                console.HtmlToAnsi("0ee6c2", ""), 
                message))
            RestoreStd()
        }
        case "error" {
            print("{}[Caster] {}ERROR:FATAL -> {}".Format(
                console.HtmlToAnsi("fa5b05", ""),
                console.HtmlToAnsi("b00000", ""),
                message))
            RestoreStd()

        }
        default {
            print("{}[Caster] Message: {}".Format(
                console.HtmlToAnsi("fa5b05", ""),
                message))
        }
    }   
}

define Help(kind="arguments") {
    VarsNonMod::HelpmapBaseFunc[kind]()
}