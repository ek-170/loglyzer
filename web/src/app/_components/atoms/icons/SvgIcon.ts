type ViewBoxValue =
  | `${number} ${number} ${number} ${number}`
  | `${number},${number},${number},${number}`;

export type SvgIconProps = {
  width?: number;
  height?: number;
  viewBox?: ViewBoxValue;
  color?: string; // not Tailwind color, use css color values
};
