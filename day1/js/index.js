import fs from "node:fs";

const input = fs.readFileSync("./input.txt", { encoding: "utf-8" });
const lines = input
  .split("\n")
  .map((line) => line.trim())
  .filter((line) => line);

let zeros = 0;
let dial = 50;

for (let i = 0; i < lines.length; i++) {
  const instruction = lines[i];
  const direction = instruction[0];
  const turns = Number(instruction.slice(1));

  for (let j = 0; j < turns; j++) {
    if (direction === "L") {
      dial--;
    } else {
      dial++;
    }

    if (dial === -1) {
      dial = 99;
    } else if (dial === 100) {
      dial = 0;
    }
  }

  if (dial === 0) {
    zeros++;
  }
}

console.log(zeros);
