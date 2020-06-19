import { render } from "https://cdn.pika.dev/@bonniernews/md2html@^0.0.8";

function read(path: string) {
  return render(Deno.readTextFileSync(path)).replace(/\n/g, "");
}
function save(path: string, text: string) {
  const encoder = new TextEncoder();
  const data = encoder.encode(text);
  Deno.writeFileSync(path, data);
}

function load(path: string, data: string[]) {
  for (const d of Deno.readDirSync(path)) {
    if (d.isDirectory) {
      // 递归
      load(path + "/" + d.name, data);
      continue;
    }
    if (!d.name.endsWith(".md")) {
      continue;
    }
    if (d.name.startsWith("README")) {
      continue;
    }
    data.push(
      `${d.name.split(".")[0]};${read(path + "/" + d.name)}`,
    );
  }
}

function main(path: string) {
  const data: string[] = [];
  load(path, data);
  save("mode.txt", data.join("\n"));
}

main("mode");
