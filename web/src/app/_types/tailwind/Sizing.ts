type VH = 'screen' | 'svh' | 'lvh' | 'dvh';

type Content = 'min' | 'max' | 'fit';

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

type Ratio =
  | '1/2'
  | '1/3'
  | '2/3'
  | '1/4'
  | '2/4'
  | '3/4'
  | '1/5'
  | '2/5'
  | '3/5'
  | '4/5'
  | '1/6'
  | '2/6'
  | '3/6'
  | '4/6'
  | '5/6'
  | 'full';

export type Width =
  | `w-${Size}`
  | `w-${Auto}`
  | `w-${VH}`
  | `w-${Ratio}`
  | `w-[${number}px]`
  | `w-${Content}`;

export type Height =
  | `h-${Size}`
  | `h-${Auto}`
  | `h-${VH}`
  | `h-${Ratio}`
  | `h-[${number}px]`
  | `h-${Content}`;
