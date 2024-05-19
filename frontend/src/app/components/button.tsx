'use client';



import Link from 'next/link'
import { useRouter } from 'next/navigation';


export function Button({ title, path }) {
    const router = useRouter();

     return (
        <button
            onClick={() => router.push(path)}
            className="
            mx-2
            bg-transparent
            hover:bg-blue-500
            text-blue-700
            font-semibold
            hover:text-white
            py-2 px-4
            border
            border-blue-500
            hover:border-transparent
            rounded"
        >
            {title}
        </button>
    )
}