import { create } from "zustand";

const useStore = create((set) => (
    {
        authName: '',
        updateAuthName: (name) => set({ authName: name }),
    }
))