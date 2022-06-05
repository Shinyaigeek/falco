import { getVCLJSON } from "./falco-integration/getVCLJSON";
import { traverse } from "./ast/traverse";
async function main() {
  // TODO
  const targetFile = "../../examples/default01.vcl";
  const vclAST = await getVCLJSON(targetFile);

  traverse(vclAST, {
    backend: async (statement) => {
      console.log("----------");
      console.log(statement.Token);
      console.log(statement.Properties);
    },
    subroutine: async (statement) => {
      console.log("----------");
      console.log(statement.Token);
      console.log(statement.Block);
    },
  });
}

main();
