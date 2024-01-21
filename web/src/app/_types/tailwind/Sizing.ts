import { Size, Auto, Ratio } from "./Unit";

type VH = 'screen' | 'svh' | 'lvh' | 'dvh';

type Content = 'min' | 'max' | 'fit';

export type Width =
  | `w-${Size}`
  | `w-${Auto}`
  | `w-${VH}`
  | `w-${Ratio}`
  | `w-${Content}`;
