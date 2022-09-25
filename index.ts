import express from "express";
import * as http from "http";

const app = express();

const server = http.createServer(app);
const port = process.env.PORT || 8000;

app.get("/", (req, res) => {
  res.json({ message: "ok" });
});

server.listen(port);
