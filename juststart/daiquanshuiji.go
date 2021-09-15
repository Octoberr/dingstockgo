public static T[] SelectWeightedItems(IEnumerable<T> ie, int count, Func<T, double> weightfunc)
{
SortedList<double, T> sd = new SortedList<double, T>();
foreach (T item in ie)
{
double weight = weightfunc(item);
if (weight <= 0) continue;
double key = Math.Pow(random.NextDouble(), (1.0 / weight));
if (sd.Count < count)
sd.Add(key, item);
else
{
sd.RemoveAt(0);
sd.Add(key, item);
}
}
return sd.Values.ToArray();
}
}