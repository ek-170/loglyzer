import type { Meta, StoryObj } from '@storybook/react';

import { Radio } from '@/app/_components/atoms/inputs/Radio';

const meta: Meta<typeof Radio> = {
  component: Radio,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof Radio>;

export const SampleRadio: Story = {
  args: {
    label: 'radio',
    color: 'light',
    name: 'sample-radio',
    radioItems: [
      {
        text: 'option 1',
      },
      {
        text: 'option 2',
      },
      {
        text: 'option 3',
      },
    ]
  },
};
