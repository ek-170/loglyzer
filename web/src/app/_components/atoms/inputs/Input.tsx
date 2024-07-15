/* eslint-disable tailwindcss/no-custom-classname */
import { Width } from '@/app/_types/tailwind/Sizing';
import { Margin, Padding } from '@/app/_types/tailwind/Spacing';
import { ComponentPropsWithRef } from 'react';
import { tv } from 'tailwind-variants';

/**
 * Button, Radio, CheckBox, File exist as another component.
 */
type InputType =
  | 'color'
  | 'date'
  | 'datetime-local'
  | 'email'
  | 'hidden'
  | 'image'
  | 'month'
  | 'range'
  | 'reset'
  | 'search'
  | 'tel'
  | 'text'
  | 'time'
  | 'url'
  | 'week';

/**
 * input "value" attr must be state created by useState
 * and input onChange must be setState function
 */
export type InputProps = {
  type: InputType;
  color: 'light' | 'dark' | 'primary';
  label?: string;
  padding?: Padding[];
  margin?: Margin[];
  width?: Width;
} & Omit<ComponentPropsWithRef<'input'>, 'type'>;

// type InputValidationItem = {
//     validator: (value: string | number | Date | null) => boolean;
//     errorMessage: string;
// };

// export type InputValidation = {
//   isValid: boolean,
//   validations: InputValidationItem[]
// };

export const Input = (props: InputProps) => {
  const { color, padding, margin, width, label, ...inputElementProps } = props;

  const paddings = padding && padding.length > 0 ? padding : ['p-0.5'];
  const margins = margin && margin.length > 0 ? margin : ['my-1.5'];
  const paddingClass = paddings.join(' ');
  const marginClass = margins.join(' ');
  const widthClass = width || 'w-fit';
  const wrapperTv = tv({
    base: `flex-col ${paddingClass} ${marginClass} ${widthClass}`,
  });

  const inputTv = tv({
    base: `peer mb-1.5 rounded border border-gray-300 invalid:border-pink-500 invalid:text-pink-600 focus:border-sky-500 focus:outline-none focus:ring-1 focus:ring-sky-500 focus:invalid:border-pink-500 focus:invalid:ring-pink-500`,
    variants: {
      color: {
        light: 'bg-black-0 text-black-600 disabled:bg-black-500',
        dark: 'bg-black-500 text-black-50 disabled:bg-black-200',
        primary: 'bg-primary-100 text-black-500 disabled:bg-primary-500',
      },
    },
  });
  return (
    <div className={wrapperTv()}>
      {label && (
        <label className="mb-2.5 block text-sm font-medium text-black-500">
          {label}
        </label>
      )}
      <input {...inputElementProps} className={inputTv({ color: color })} />
      {inputElementProps.type === 'url' ? (
        <p className="invisible text-sm text-pink-500 peer-invalid:visible">
          Please enter in Valid URL
        </p>
      ) : inputElementProps.type === 'email' ? (
        <p className="invisible text-sm text-pink-500 peer-invalid:visible">
          Please enter in Valid Email
        </p>
      ) : inputElementProps.type === 'tel' ? (
        <p className="invisible text-sm text-pink-500 peer-invalid:visible">
          Please enter in Valid Telephone Number
        </p>
      ) : (
        <p className="invisible text-sm text-pink-500 peer-invalid:visible">
          Please enter in Valid value
        </p>
      )}
    </div>
  );
};
