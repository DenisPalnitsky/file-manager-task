'use client'

import { usePathname } from 'next/navigation'
import Link from 'next/link'
import './globals.css'


export function Navigation() {
    const pathname = usePathname()
   
    return (
  
      <nav className="bg-blue-500 p-4"> 
      <div className="flex justify-between items-center max-w-4xl mx-auto text-white">      
      <Link className="text-white hover:text-gray-200" href="/">Home</Link>
        <Link className="text-white hover:text-gray-200" href="/upload">Upload file</Link>
        <Link className="text-white hover:text-gray-200" href="/list">List files</Link>       
      </div>
    </nav>
  
    )
  }