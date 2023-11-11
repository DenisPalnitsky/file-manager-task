
export const revalidate = 0


async function fetchFiles() {
    const res = await fetch(`http://localhost:8080/files`, { cache: 'no-store' })
    if (!res.ok) {
        // This will activate the closest `error.js` Error Boundary
        throw new Error('Failed to fetch data')
    }

    return res.json()
}


export default async function FileList() {
    const files = await fetchFiles()

    return (<table className="table-auto border-collapse border border-gray-400">
        <thead>
            <tr>
                <th className="px-4 py-2 bg-gray-200 text-left text-xs font-semibold uppercase border-b border-gray-400">Filename</th>
            </tr>
        </thead>
        <tbody>
            {files.files.map((file: { Filename: string }) => (
                <tr key={file.Filename}>
                    <td className="border px-4 py-2">{file.Filename}</td>
                </tr>
            ))}
        </tbody>
    </table>)
}