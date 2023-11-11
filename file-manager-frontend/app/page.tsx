'use client'

import { useState } from 'react';

export default function Home() {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [uploadResponse, setUploadResponse] = useState<string | null>(null);
  const [uploading, setUploading] = useState<boolean>(false);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      setUploadResponse(``);
      setSelectedFile(event.target.files[0]);      
    }
  };

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (!selectedFile) {
      return;
    }
    const formData = new FormData();
    formData.append('file', selectedFile);
    try {
      setUploading(true);
      const response = await fetch('http://localhost:8080/files', {
        method: 'POST',
        body: formData,
      });
      if (response.ok) {
        console.log('File uploaded successfully');
        setUploadResponse('File uploaded successfully');        
      } else {
        console.error('Failed to upload file');
      }
    } catch (error) {
      setUploadResponse(`Upload failed: ${error}`);
      console.error('Failed to upload file', error);
    } finally {
      setUploading(false);
    }
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24 bg-gray-100">
      <div className="bg-white rounded-lg shadow-lg p-8 items-center flex flex-col items-center">
        <h1 className="text-4xl font-bold mb-4">File Manager</h1>
        <p className="text-lg mb-4">
          Upload your files <a href="/upload" className="text-blue-500 hover:underline">here</a> and get a list of your files <a href="/list" className="text-blue-500 hover:underline">here</a>.
        </p>        
      </div>
    </main>
  );
}
