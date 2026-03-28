import fs from "node:fs";

const file = fs.readFileSync("../input.txt", { encoding: "utf-8" }).trim();
const idRanges = file.split(",");

const invalidIds = [];

idRanges.forEach((range) => {
  const [start, end] = range.split("-");

  const startNum = Number(start);
  const endNum = Number(end);

  for (let i = startNum; i <= endNum; i++) {
    const strId = String(i);

    if (strId.length % 2 !== 0) {
      continue;
    }

    const firstHalf = strId.slice(0, strId.length / 2);
    const secondHalf = strId.slice(strId.length / 2);

    if (firstHalf === secondHalf) {
      invalidIds.push(Number(strId));
    }
  }
});

const sum = invalidIds.reduce((acc, curr) => acc + curr, 0);
console.log(sum);
