
define VerifyHostOnNetwork(addr) {
    set Res := json.Parse("Backend/UnitDatabase/Capture.json")["Device"];
    foreach section in Res {
        
        if (section["IPA"] == addr) {
            println("Found address with info:")
            println(" \t MAC  | {}".Format(section["MAC"]))
            println(" \t Name | {}".Format(section["Name"]))
            set newdata := `{
    "Device": [
        {
            "IP": "{}",
            "Brand": "{}",
            "Database": "{}" 
        }
    ]
}`.Format(
    section["IPA"],
    section["Name"],
    section["SkyLine_DB"]);
            File.New("Backend/UnitDatabase/TargetInfo.json")
            File.Overwrite(newdata)
        }
    };
}


