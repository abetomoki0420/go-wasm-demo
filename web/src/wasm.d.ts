declare global {
  interface Window {
    formatJSON: (input: string) => string;
    add: (a: number, b: number) => number;
    sum: (a: number[]) => number;
  }
}

export {}