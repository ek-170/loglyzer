import type { Meta, StoryObj } from '@storybook/react';

import { CheckBox } from '@/app/_components/atoms/inputs/CheckBox';

const meta: Meta<typeof CheckBox> = {
  component: CheckBox,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof CheckBox>;

export const SampleCheckBox: Story = {
  args: {
    label: 'checkbox',
    color: 'light',
    text: 'this is Check Box',
  },
};
