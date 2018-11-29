const fs = require("fs");
const json2toml = require("json2toml");
const toki = require("./tokipona.json");

var result = {};
var kinds = {};

toki.forEach((koha) => {
    let name = koha["name"];
    delete koha["name"];
    result[name] = koha;

    koha.grammar.forEach((kohe) => {
        if (kinds[kohe] == null) {
            kinds[kohe] = 0;
        }

        kinds[kohe] += 1;
    });
});

fs.writeFile("tokipona.toml", json2toml(result), (err) => {
    if (err != null) {
        throw err;
    }
});

console.log(kinds);
