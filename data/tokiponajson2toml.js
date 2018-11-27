const fs = require("fs");
const json2toml = require("json2toml");
const toki = require("./tokipona.json");

var result = {};

toki.forEach((obj) => {
    result[obj["name"]] = obj;
});

fs.writeFile("tokipona.toml", json2toml(result), (err) => {
    if (err != null) {
        throw err;
    }
});
