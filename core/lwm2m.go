package core

type TypeCode        int
type OperationCode   int

const (
    TYPE_STRING   TypeCode = 0
    TYPE_INTEGER  TypeCode = 1
    TYPE_FLOAT    TypeCode = 2
    TYPE_BOOLEAN  TypeCode = 3
    TYPE_OPAQUE   TypeCode = 4
    TYPE_TIME     TypeCode = 5
)

const (
    OPERATION_NONE  OperationCode = 0
    OPERATION_R     OperationCode = 1
    OPERATION_W     OperationCode = 2
    OPERATION_RW    OperationCode = 3
    OPERATION_E     OperationCode = 4
    OPERATION_RE    OperationCode = 5
    OPERATION_WE    OperationCode = 6
    OPERATION_RWE   OperationCode = 7
)