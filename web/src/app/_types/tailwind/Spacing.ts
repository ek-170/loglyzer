import { Size } from './Unit';

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
