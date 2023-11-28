const { parentPort } = require("worker_threads");
const JSZip = require("jszip");
const fs = require("fs");
const path = require("path");

parentPort.on("message", async () => {
  const imagesDirectory = path.join(__dirname, "images");
  const zip = new JSZip();

  fs.readdirSync(imagesDirectory).forEach((file) => {
    const filePath = path.join(imagesDirectory, file);
    const data = fs.readFileSync(filePath);
    zip.file(file, data);
  });

  const zipData = await zip.generateAsync({ type: "nodebuffer" });
  fs.writeFileSync(path.join(__dirname, "output.zip"), zipData);
  parentPort.postMessage("File zip đã được tạo: output.zip");
});
