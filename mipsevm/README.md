# `mipsevm`

Supported 55 instructions:
```
'addi', 'addiu', 'addu', 'and', 'andi',
'b', 'beq', 'beqz', 'bgez', 'bgtz', 'blez', 'bltz', 'bne', 'bnez',
'clz', 'divu',
'j', 'jal', 'jalr', 'jr',
'lb', 'lbu', 'lui', 'lw', 'lwr',
'mfhi', 'mflo', 'move', 'movn', 'movz', 'mtlo', 'mul', 'multu',
'negu', 'nop', 'not', 'or', 'ori',
'sb', 'sll', 'sllv', 'slt', 'slti', 'sltiu', 'sltu', 'sra', 'srl', 'srlv', 'subu', 'sw', 'swr', 'sync', 'syscall',
'xor', 'xori'
```

To run:
1. Load a program into a state, e.g. using `LoadELF`.
2. Patch the program if necessary: e.g. using `PatchVM` for Go programs.
3. Load the state into a MIPS-32 configured unicorn instance, using `NewUnicorn`, `LoadUnicorn`
4. Implement the `PreimageOracle` interface
5. Instrument the emulator with the state, and pre-image oracle, using `NewUnicornState`
6. Step through the instrumented state with `Step(proof)`,
   where `proof==true` if witness data should be generated. Steps are faster with `proof==false`.
7. Optionally repeat the step on-chain by calling `MIPS.sol` and `Oracle.sol`, using the above witness data.