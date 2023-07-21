import "./globals.css";

import { CssBaseline } from "@mui/material";
import { Inter } from "next/font/google";

import { WNThemeProvider } from "@/components/atoms/WNThemeProvider";
import { WNLayout } from "@/components/layout/WNLayout";
import { Metadata } from "next";

const inter = Inter({ subsets: ["latin"] });
export const metadata: Metadata = {
  title: "Winnin Tournaments",
  description: "Winnin is a tournament hosting platform",
};
export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <WNThemeProvider>
      <CssBaseline>
        <html lang="en">
          <body className={inter.className}>
            <WNLayout>{children}</WNLayout>
          </body>
        </html>
      </CssBaseline>
    </WNThemeProvider>
  );
}
