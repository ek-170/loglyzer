type PaddingDirection =
  | 'p'
  | 'px'
  | 'py'
  | 'ps'
  | 'pe'
  | 'pt'
  | 'pr'
  | 'pb'
  | 'pl';

export type Padding = `${PaddingDirection}-${Size}`;

type MarginDirection =
  | 'm'
  | 'mx'
  | 'my'
  | 'ms'
  | 'me'
  | 'mt'
  | 'mr'
  | 'mb'
  | 'ml';

export type Margin = `${MarginDirection}-${Size}`;

type Size =
  | 'px'
  | 0
  | 0.5
  | 1
  | 1.5
  | 2
  | 2.5
  | 3
  | 3.5
  | 4
  | 5
  | 6
  | 7
  | 8
  | 9
  | 10
  | 11
  | 12
  | 14
  | 16
  | 20
  | 24
  | 28
  | 32
  | 36
  | 40
  | 44
  | 48
  | 52
  | 56
  | 60
  | 64
  | 72
  | 80
  | 96;
