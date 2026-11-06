// JS 沒有 interface:「工廠」就是任何有 createButton / createCheckbox 的物件(duck typing)。

export const darkFactory = {
  createButton: () => ({ render: () => '[dark button]' }),
  createCheckbox: () => ({ render: () => '[dark checkbox]' }),
};

export const lightFactory = {
  createButton: () => ({ render: () => '[light button]' }),
  createCheckbox: () => ({ render: () => '[light checkbox]' }),
};

export function renderUI(factory) {
  return `${factory.createButton().render()} ${factory.createCheckbox().render()}`;
}
