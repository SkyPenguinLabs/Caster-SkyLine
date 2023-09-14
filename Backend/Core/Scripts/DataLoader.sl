set TableColumns := ["Device Name", "MAC", "Address"];

set Rows := [];

console.TableNew(
    "", "", "", "", "",
    `━`, `┃`, `┫`, `┣`);

define FileLoader(filename){
    return json.Parse(filename);
};


define TabulateHosts(devicetypeappend="all") {
    print(" \n ")
    set Res := FileLoader("Backend/UnitDatabase/Capture.json")["Device"];
    foreach section in Res {
        if (devicetypeappend == "all") {
            set newarr := [section["Name"], section["MAC"], section["IPA"]];
            Rows.Append(newarr);
        } else {
            if(section["Name"].StartsWith(devicetypeappend) == true) {
                set newarr := [section["Name"], section["MAC"], section["IPA"]];
                Rows.Append(newarr);
            }
        }
    };
    println(console.Table(Rows, TableColumns))
    Rows.Reset()
    io.restore()
}

