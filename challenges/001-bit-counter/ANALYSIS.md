# Analysis: Bit Counter

## Approach & Invariants

So the whole exercise was to sharpen the skills in using bitwise operators, since they are faster than the regular operations like `+,-,/,*`.

The approach is easy - 
1. We get the input number and compare it with `AND or &` by binary representation of `1<<i`
2. We iterate through each bit position (0 to 31 for 32-bit integers, or 0 to 63 for 64-bit)
3. For each position `i`, we create a mask using left shift: `1 << i` (this gives us a number with only the i-th bit set)
4. We perform bitwise AND between our input number and the mask: `number & (1 << i)`
5. If the result is non-zero, it means the i-th bit is set to 1, so we increment our counter
6. Continue until we've checked all bit positions

**Invariants:**
- The bit counter never decreases during iteration
- Each bit position is checked exactly once
- The mask `1 << i` always has exactly one bit set at position `i`

## Correctness Argument

**Claim:** The algorithm correctly counts all set bits (1s) in the binary representation of any non-negative integer.

**Proof sketch:**
1. **Completeness**: We examine every bit position from 0 to (word_size - 1), ensuring no set bits are missed.
2. **Accuracy**: For each position `i`, the mask `1 << i` has exactly one bit set at position `i`. When we compute `number & (1 << i)`:
   - If bit `i` in `number` is 1: `1 & 1 = 1` (non-zero result, increment counter)
   - If bit `i` in `number` is 0: `0 & 1 = 0` (zero result, no increment)
3. **No double counting**: Each bit position is tested exactly once, so no bit contributes more than once to the final count.

## Time & Space Complexity

### Time Complexity
- **Big-O**: O(log n) where n is the input number, or O(w) where w is the word size (32 or 64 bits)
- **Practical Notes**: For fixed-width integers, this becomes O(1) since we always check a constant number of bits (32 or 64). The loop runs exactly w times regardless of input value.

### Space Complexity
- **Big-O**: O(1)
- **Practical Notes**: Uses only a constant amount of extra memory for the counter variable and loop index, regardless of input size.

## Test Plan

### Test Cases Covered
- **Zero**: `0` → should return `0` (no bits set)
- **Powers of 2**: `1, 2, 4, 8, 16, ...` → should return `1` (exactly one bit set)
- **All bits set**: `255` (for 8-bit), `65535` (for 16-bit) → should return `8`, `16` respectively
- **Mixed patterns**: `7` (binary: 111) → should return `3`
- **Large numbers**: Test with maximum values for the data type
- **Edge cases**: `1` (minimum positive), maximum integer value