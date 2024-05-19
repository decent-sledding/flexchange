'use client';

import { SubmitButton, LabeledInput } from '@/app/components/form';


async function subscribeWithEmail() {
    const form = document.querySelector("#subscription_form") as HTMLFormElement;
    const apiUrl = process.env['BACKEND_URL'];

    if (!form) {
        console.error("Could not find HTML form for email subscription");
        return ;
    }

    if (apiUrl === undefined) {
        console.error("Backend URL is not set!");
        alert("Something is wrong! Need to check forntend settings");
        return ;
    }

    console.log(`Using server URL: ${apiUrl}`);

    const response = await fetch(String(apiUrl), {
        method: 'POST',
        headers: {
            "Content-Type": "application/json",
        },
        body: new FormData(form),
    });

    if (!response.ok) {
        alert("Could not subscribe");
        return ;
    }

    alert("Successfully subscribed!");
}

function handleSubmit(e: Event) {
    console.log("Making request to the backend");
    e.preventDefault();
    subscribeWithEmail().then(() => {});
}

export default function Page() {
    return (
        <main>
            <div className="flex justify-center py-5 text-black	">
                <h1 className="header text-lg">Rate Update Email Subscription</h1>
            </div>
            <div className="flex max-w-full w-full h-screen justify-center">
                <div className="m-auto">
                    <form onSubmit={handleSubmit} id="subscription_form" className="w-full">
                        <LabeledInput
                            inputId="first"
                            inputType="text"
                            inputName="email"
                            labelText="Your E-Mail"/>
                        <SubmitButton formId="subscription_form" buttonText="Subscribe"/>
                    </form>
                </div>
            </div>
        </main>
    );
}