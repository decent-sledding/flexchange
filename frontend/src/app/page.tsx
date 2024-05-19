import Image from "next/image";
import { Button } from './components/button';

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-32">
      <div className="inline">
        <Button title="See Rates" path='/services/rate' />
        <Button title="Subscription" path='/services/subscribe'/>
      </div>
    </main>
  );
}
