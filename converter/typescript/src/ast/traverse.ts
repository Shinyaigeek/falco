import type { StatementType, VCLAST } from "./type";

type Traverser = {
  backend?: (statement: StatementType) => Promise<void>;
  subroutine?: (statement: StatementType) => Promise<void>;
};

export const traverse: (ast: VCLAST, traverser: Traverser) => Promise<void> =
  async function (ast, traverser) {
    for (const statement of ast) {
      switch (statement.Token.Type) {
        case "BACKEND": {
          if (traverser.backend) {
            await traverser.backend(statement);
            break;
          }
        }
        case "SUBROUTINE": {
          if (traverser.subroutine) {
            await traverser.subroutine(statement);
            break;
          }
        }
      }
    }
  };
