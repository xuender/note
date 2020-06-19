import { render } from "https://cdn.pika.dev/@bonniernews/md2html@^0.0.8";

function read(path: string) {
  return render(Deno.readTextFileSync(path)).replace(/\n/g, "");
}
function save(path: string, text: string) {
  const encoder = new TextEncoder();
  const data = encoder.encode(text);
  Deno.writeFileSync(path, data);
}

function main(path: string) {
  const data = [];
  for (const dirEntry of Deno.readDirSync(path)) {
    if (!dirEntry.isDirectory) {
      continue;
    }
    for (const d of Deno.readDirSync(path + "/" + dirEntry.name)) {
      if (d.isDirectory || !d.name.endsWith(".md")) {
        continue;
      }
      data.push(
        d.name.split(".")[0] + ";" +
          read(path + "/" + dirEntry.name + "/" + d.name),
      );
    }
  }
  save("dict.txt", data.join("\n"));
}

main("../mode");
