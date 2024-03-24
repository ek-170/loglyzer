import { SvgIconProps } from './SvgIcon';

export const CrossIcon = (props: SvgIconProps) => {
  const viewBox = props.viewBox || '0 0 512 512';
  const width = props.width || 30;
  const height = props.height || 30;
  const color = props.color || 'rgb(75, 75, 75)';
  return (
    <div className="grid place-items-center">
      <svg
        version="1.1"
        id="_x32_"
        role="img"
        xmlns="http://www.w3.org/2000/svg"
        viewBox={viewBox}
        width={`${width}px`}
        height={`${height}px`}
        opacity="1"
      >
        <g>
          <polygon
            points="511.998,70.682 441.315,0 256.002,185.313 70.685,0 0.002,70.692 185.316,256.006 0.002,441.318
		70.69,512 256.002,326.688 441.315,512 511.998,441.318 326.684,256.006"
            fill={color}
          ></polygon>
        </g>
      </svg>
    </div>
  );
};
