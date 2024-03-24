import { ReactElement } from 'react';
import { render, RenderOptions } from '@testing-library/react';
import userEvent, { Options } from '@testing-library/user-event';

function setup(
  ui: ReactElement,
  renderOptions?: RenderOptions,
  options?: Options,
) {
  return {
    ...render(ui, {
      // wrapper: yourProvider,
      ...renderOptions,
    }),
    user: userEvent.setup(options),
  };
}

export * from '@testing-library/react';
export { setup as render };
