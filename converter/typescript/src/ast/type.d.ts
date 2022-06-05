// TODO: 真面目に型定義する
export type VCLAST = StatementType[];

type TokenType =
  | "ILLEGAL"
  | "EOF"
  | "IDENT"
  | "INT"
  | "STRING"
  | "FLOAT"
  | "RTIME"
  | "COMMENT"
  | "TRUE"
  | "FALSE"
  | "LF"
  | "EQUAL"
  | "NOTEQUAL"
  | "REGEX"
  | "NOT_REGEX_MATCH"
  | "GREATER_THAN"
  | "LESS_THAN"
  | "GREATER_THAN_EQUAL"
  | "LESS_THAN_EQUAL"
  | "AND"
  | "OR"
  | "ASSIGN"
  | "ADDITION"
  | "SUBTRACTION"
  | "MULTIPLICATION"
  | "DIVISION"
  | "REMAINDER"
  | "BITWISE_OR"
  | "BITWISE_AND"
  | "BITWISE_XOR"
  | "LEFT_SHIFT"
  | "RIGHT_SHIFT"
  | "LEFT_ROTATE"
  | "RIGHT_ROTATE"
  | "LOGICAL_AND"
  | "LOGICAL_OR"
  | "LEFT_BRACE"
  | "RIGHT_BRACE"
  | "LEFT_PAREN"
  | "RIGHT_PAREN"
  | "LEFT_BRACKET"
  | "RIGHT_BRACKET"
  | "COMMA"
  | "SLASH"
  | "SEMICOLON"
  | "DOT"
  | "NOT"
  | "COLON"
  | "PLUS"
  | "MINUS"
  | "ACL"
  | "DIRECTOR"
  | "BACKEND"
  | "TABLE"
  | "SUBROUTINE"
  | "ADD"
  | "CALL"
  | "DECLARE"
  | "ERROR"
  | "ESI"
  | "INCLUDE"
  | "IMPORT"
  | "LOG"
  | "REMOVE"
  | "RESTART"
  | "RETURN"
  | "SET"
  | "SYNTHETIC"
  | "SYNTHETIC_BASE64"
  | "UNSET"
  | "IF"
  | "ELSE"
  | "ELSEIF"
  | "ELSIF"
  | "PENALTYBOX"
  | "RATECOUNTER";

// TODO: define only necessary, but all property should be defined
type Token = {
  Type: TokenType;
  Literal: string;
};

type TableDeclarationType = {
  Token: Token;
  Name: NameType;
  Properties: TablePropertyType[];
  Block: BlockType
};

type TablePropertyType = {
  Token: Token;
  Key: string;
  Value: ExpressionType;
};

type ExpressionType = {} & NodeType;

type StatementType = {} & NodeType & TableDeclarationType;

type NameType = {} & NodeType;

type BlockType = {
    Statements: StatementType[]
} & NodeType;

type NodeType = {
  Token: Token;
  Value: string;
};
