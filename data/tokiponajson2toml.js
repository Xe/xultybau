const fs = require("fs");
const json2toml = require("json2toml");
const toki = require("./tokipona.json");

var result = {};
var kinds = {};

toki.forEach((obj) => {
    let name = obj["name"];
    delete obj["name"];
    result[name] = obj;

    obj.grammar.forEach((koha) => {
        if (kinds[koha] == null) {
            kinds[koha] = 0;
        }

        kinds[koha] += 1;
    });
});

fs.writeFile("tokipona.toml", json2toml(result), (err) => {
    if (err != null) {
        throw err;
    }
});

console.log(kinds);
