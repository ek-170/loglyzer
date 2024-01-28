import type { Meta, StoryObj } from '@storybook/react';

import { Select } from '@/app/_components/atoms/inputs/Select';

const meta: Meta<typeof Select> = {
  component: Select,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof Select>;

export const SampleSelect: Story = {
  args: {
    label: 'radio',
    color: 'light',
    name: 'sample-radio',
    optionItems: [
      {
        text: 'list 1',
        value: 'list_1',
      },
      {
        text: 'list 2',
        value: 'list_2',
      },
      {
        text: 'list 3',
        value: 'list_3',
      },
    ],
  },
};
