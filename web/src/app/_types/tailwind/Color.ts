type DefaultColorName =
  | 'primary'
  | 'black'
  | 'white'
  | 'gray'
  | 'red'
  // | "orange"
  // | "amber"
  // | "yellow"
  // | "lime"
  // | "green"
  // | "emerald"
  // | "teal"
  // | "cyan"
  | 'sky'
  | 'blue'
  | 'indigo'
  | 'violet'
  | 'purple'
  | 'fuchsia'
  | 'pink'
  | 'rose';

type Depth = 50 | 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | 950;

export type DefaultColor = `${DefaultColorName}-${Depth}`;