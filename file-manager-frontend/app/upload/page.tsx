'use client'

import { useState } from 'react';

export default function Page() {
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
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <form onSubmit={handleSubmit} className="flex flex-col items-center justify-center h-full">
          <div className="mb-4">
            <label htmlFor="file" className="block text-gray-700 font-bold mb-2">
              Select file to upload
            </label>            
          </div>
          <div className='mb-4'>
            <input type="file" onChange={handleFileChange} className="appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" />
          </div>
          <div className="flex justify-center">
            <button type="submit" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
              Upload
            </button>
          </div>
          <div className="flex justify-center">
            {uploading && <p>Uploading...</p>}
            {uploadResponse && !uploading && <p>{uploadResponse}</p>}
          </div>
        </form>
      </main>
  );
}
