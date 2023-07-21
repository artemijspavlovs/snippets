"use client";

import { WNThemeGenerate } from "./mui";
import { PaletteMode, ThemeProvider, createTheme } from "@mui/material";
import { PropsWithChildren, useMemo, useState } from "react";
import { ColorThemeContext } from "./ThemeToggleButton";

export const WNThemeProvider = ({ children }: PropsWithChildren) => {
  const [mode, setMode] = useState<PaletteMode>("light");
  const theme = useMemo(() => createTheme(WNThemeGenerate(mode)), [mode]);
  const colorMode = useMemo(
    () => ({
      toggleColorMode: () => {
        setMode((prevMode: PaletteMode) =>
          prevMode === "light" ? "dark" : "light",
        );
      },
    }),
    [],
  );

  return (
    <ColorThemeContext.Provider value={colorMode}>
      <ThemeProvider theme={theme}>{children}</ThemeProvider>
    </ColorThemeContext.Provider>
  );
};
