import React, { useEffect, useState } from 'react';
import FileList  from './file-list';


export default async function Page() {
    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
          <FileList />
        </main>
    )
}