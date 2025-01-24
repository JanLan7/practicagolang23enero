#programa para temperatura
temperatura = int(input(" Ingresa la temperatura en celcius: "))

if temperatura > 27:
    print("Comprar helado")
elif temperatura < 15:
    print("Comprar chocolate")
else:
    print("Comprar jugo de naranja")

print("Fin del programa")
