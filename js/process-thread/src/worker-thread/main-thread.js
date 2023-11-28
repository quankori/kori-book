const { Worker } = require("worker_threads");

// Tạo Worker Thread thứ nhất
const worker1 = new Worker("./src/worker-thread/child-thread.js");
worker1.on("message", (message) =>
  console.log("Tin nhắn từ Worker 1:", message)
);
worker1.postMessage("Dữ liệu cho Worker 1");

// Tạo Worker Thread thứ hai
const worker2 = new Worker("./src/worker-thread/child-thread.js");
worker2.on("message", (message) =>
  console.log("Tin nhắn từ Worker 2:", message)
);
worker2.postMessage("Dữ liệu cho Worker 2");
