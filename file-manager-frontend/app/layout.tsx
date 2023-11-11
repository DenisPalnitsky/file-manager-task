import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'
import { Navigation } from './navigation'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Create Next App',
  description: 'Generated by create next app',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">      
      <body className={inter.className}>
        <header className="text-3xl font-bold underline text-center py-4 bg-blue-500 text-white">File manager</header>
        <Navigation />
        {children}
      </body>
    </html>
  )
}
