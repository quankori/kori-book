const { spawn } = require("child_process");

const child = spawn("sleep", ["5"]);

child.on("exit", function (code, signal) {
  console.log("Child process exited with code", code);
});

// Main process ghi ra console ngay lập tức
console.log("Main process logged this message first.");
