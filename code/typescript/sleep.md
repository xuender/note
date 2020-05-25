# sleep

睡眠等待

```typescript
export function sleep(millisecond: number) {
    return new Promise(resolve => setTimeout(resolve, millisecond))
}
```
