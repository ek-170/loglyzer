export type Top = `top-${Size}` | `top-${Auto}` | `top-${Ratio}`;

export type Bottom = `bottom-${Size}` | `bottom-${Auto}` | `bottom-${Ratio}`;

export type Right = `right-${Size}` | `right-${Auto}` | `right-${Ratio}`;

export type Left = `left-${Size}` | `left-${Auto}` | `left-${Ratio}`;

export type Start = `start-${Size}` | `start-${Auto}` | `start-${Ratio}`;

export type Inset =
  | `inset-${Size}`
  | `inset-${Auto}`
  | `inset-${Ratio}`
  | `inset-x-${Size}`
  | `inset-x-${Auto}`
  | `inset-x-${Ratio}`
  | `inset-y-${Size}`
  | `inset-y-${Auto}`
  | `inset-y-${Ratio}`;

export type End = `end-${Size}` | `end-${Auto}` | `end-${Ratio}`;

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

type Auto = 'auto';

type Ratio = '1/2' | '1/3' | '2/3' | '1/4' | '2/4' | '3/4' | 'full';
