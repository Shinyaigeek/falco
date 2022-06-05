import { getVCLJSON } from "./falco-integration/getVCLJSON";
async function main() {
  // TODO
  const targetFile = "../../examples/default01.vcl";
  console.dir(await getVCLJSON(targetFile), { depth: null });
}

main();
