import { Width } from '@/app/_types/tailwind/Sizing';
import { Margin, Padding } from '@/app/_types/tailwind/Spacing';
import { ComponentPropsWithRef } from 'react';
import { tv } from 'tailwind-variants';

/**
 * Button, Radio, CheckBox, File exist as another component.
 */
type InputType =
  | "color"
  | "date"
  | "datetime-local"
  | "email"
  | "hidden"
  | "image"
  | "month"
  | "range"
  | "reset"
  | "search"
  | "tel"
  | "text"
  | "time"
  | "url"
  | "week";

/**
 * padding and width should be written in the way the class name of tailwind is written
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

  const paddings =
    padding && padding.length > 0
      ? padding
      : ['p-0.5'];
  const margins =
    margin && margin.length > 0
      ? margin
      : ['my-1.5'];
  const paddingClass = paddings.join(' ');
  const marginClass = margins.join(' ')
  const widthClass = width || 'w-fit';
  // eslint-disable-next-line tailwindcss/no-custom-classname
  const input = tv({
    base: `rounded border border-gray-300 ${paddingClass} ${marginClass} ${widthClass} peer invalid:border-pink-500 invalid:text-pink-600 focus:border-sky-500 focus:outline-none focus:ring-1 focus:ring-sky-500 focus:invalid:border-pink-500 focus:invalid:ring-pink-500`,
    variants: {
      color: {
        light: 'text-black-600 bg-black-0 disabled:bg-black-500',
        dark: 'text-black-50 bg-black-500 disabled:bg-black-200',
        primary: 'text-black-500 bg-primary-100 disabled:bg-primary-500',
      },
    },
  });
  return (
    <div className="flex-col">
      {label && <label className="block text-sm font-medium text-black-500">{label}</label>}
      <input {...inputElementProps} className={input({ color: color })} />
      {
        inputElementProps.type === "url" ?
          <p className="invisible text-sm text-red-500 peer-invalid:visible">
            Please enter in Valid URL
          </p> :
        inputElementProps.type === "email" ?
          <p className="invisible text-sm text-red-500 peer-invalid:visible">
            Please enter in Valid Email
          </p> :
        inputElementProps.type === "tel" ?
          <p className="invisible text-sm text-red-500 peer-invalid:visible">
            Please enter in Valid Telephone Number
          </p> :
          <p className="invisible text-sm text-red-500 peer-invalid:visible">
            Please enter in Valid value
          </p>
      }
    </div>
  );
};
