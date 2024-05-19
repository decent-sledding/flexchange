interface BasicInputData {
    inputId: string,
    inputType: string,
    inputValue?: string,
    inputText?: string,
    inputName: string,
}

interface BasicLabelData {
    labelText: string,
    labelForId?: string,
}

interface SubmitButtonData {
    formId: string,
    buttonText: string,
}

interface LabeledInputData extends BasicInputData, BasicLabelData {}



export function BasicLabel(props: BasicLabelData) {
    return (
        <label
            htmlFor={props.labelForId}
            className=""
        >
            {props.labelText}
        </label>
    );
}


export function BasicInput(props: BasicInputData) {
    return (
        <input
            className="
                outline-none
                bg-gray-50
                border
                border-gray
                text-gray-900
                text-sm
                rounded-lg
                focus:ring-blue-500
                focus:border-blue-500
                block
                w-full
                p-3
                my-1
                dark:bg-gray-700
                dark:border-gray-600
                dark:placeholder-gray-400
                dark:text-white
                dark:focus:ring-blue-500
                dark:focus:border-blue-500"
            id={props.inputId}
            type={props.inputType}
            name={props.inputName}
            value={props.inputValue}
        />
    );
}


export function LabeledInput(props: LabeledInputData) {
    return (<>
        <fieldset>
            <BasicLabel labelText={props.labelText} labelForId={props.inputId}/>
            <BasicInput {...props} />
        </fieldset>
    </>);
}


export function SubmitButton(props: SubmitButtonData) {
    return (
        <button
            className="
                min-w-[320px]
                bg-white
                hover:bg-gray-100
                text-gray-800
                font-semibold
                py-4
                px-4
                my-3
                border
                border-gray-400
                rounded
                shadow
                btn
                btn-green
                w-24"
            type="submit"
            form={props.formId}
        >
            {props.buttonText}
        </button>
    );
}